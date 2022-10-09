package biz

import (
	"bytes"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	json "github.com/json-iterator/go"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"strings"
	"time"
)

type RsaRepo interface {
	Fetch(string) (*RsaObj, error)
	Store(string, time.Duration, *RsaObj) error
}

type RsaUsecase struct {
	repo   RsaRepo
	logger *zap.Logger
}

type RsaHelper struct {
	repo   RsaRepo
	logger *zap.Logger
	err    error
}

func NewRsaUsecase(repo RsaRepo, logger *zap.Logger) *RsaUsecase {
	SetRsaKeyPrefix("rsa")
	return &RsaUsecase{repo: repo, logger: logger}
}

func (g *RsaUsecase) Helper() *RsaHelper {
	return &RsaHelper{repo: g.repo, logger: g.logger}
}

func (h *RsaHelper) Err() error {
	return h.err
}

func (g *RsaUsecase) PublicKey(requestId string, opts ...Option) ([]byte, error) {
	if len(requestId) == 0 {
		return nil, errors.New("empty requestId for fetch rsa key")
	}

	requestId = g.solveId(requestId)
	obj, err := g.repo.Fetch(requestId)
	if err != nil {
		return nil, err
	}

	if obj == nil {
		gen := &generator{global.bits, true, time.Duration(-1)}
		for _, opt := range opts {
			opt(gen)
		}

		if !gen.renew {
			return nil, status.Errorf(codes.NotFound, "can't find rsaObj")
		}

		gen.init()
		obj, err = gen.genRsaObj()
		if err != nil {
			return nil, err
		}

		err = g.repo.Store(requestId, gen.expiration, obj)
		if err != nil {
			return nil, err
		}
	}
	return obj.Public, nil
}

func (g *RsaUsecase) solveId(requestId string) string {
	if len(global.prefix) > 0 {
		requestId = strings.TrimPrefix(requestId, ":")
	}
	requestId = global.prefix + requestId
	return requestId
}

// RsaObj is a rsa keys instance
type RsaObj struct {
	requestId string
	Private   []byte
	Public    []byte
}

func (o *RsaObj) MarshalBinary() (data []byte, err error) {
	return json.Marshal(o)
}

func (o *RsaObj) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, o)
}

// Encrypt returns to encode data and if an error
func (o *RsaObj) Encrypt(data []byte) ([]byte, error) {
	block, _ := pem.Decode(o.Public)
	if block == nil {
		return nil, errors.New("public key error")
	}
	pubKey, err := x509.ParsePKCS1PublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	return rsa.EncryptPKCS1v15(rand.Reader, pubKey, data)
}

// Decrypt returns to decode data and if an error
func (o *RsaObj) Decrypt(ciphertext []byte) ([]byte, error) {
	block, _ := pem.Decode(o.Private)
	if block == nil {
		return nil, errors.New("private key error")
	}
	private, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	return rsa.DecryptPKCS1v15(rand.Reader, private, ciphertext)
}

// globalOption is using default option
type globalOption struct {
	prefix     string
	bits       int
	expiration time.Duration
}

func SetRsaBits(bits int) {
	global.bits = bits
}

func SetRsaExpiration(expiration time.Duration) {
	global.expiration = expiration
}

// SetRsaKeyPrefix settings the redis cached rsa key prefix
func SetRsaKeyPrefix(prefix string) {
	if len(prefix) > 0 && !strings.HasSuffix(prefix, ":") {
		prefix += ":"
	}
	global.prefix = prefix
}

type Option func(*generator)

func WithBits(bits int) Option {
	return func(g *generator) {
		g.bits = bits
	}
}

// WithExpiration settings the global rsa expiration
// Zero expiration means the key has no expiration time.
func WithExpiration(expiration time.Duration) Option {
	return func(g *generator) {
		g.expiration = expiration
	}
}

// WithNoGen when the requestId is not exist don't create
func WithNoGen(g *generator) {
	g.renew = false
}

var global = &globalOption{bits: 1024, expiration: time.Minute * 10}

// generator create a rsa key
type generator struct {
	bits       int
	renew      bool
	expiration time.Duration
}

func (g *generator) init() {
	if g.bits <= 0 {
		g.bits = global.bits
	}
	if g.expiration < 0 {
		g.expiration = global.expiration
	}
}

// genRsaObj generated rsa private and public key
func (g *generator) genRsaObj() (*RsaObj, error) {
	privateKey, err := rsa.GenerateKey(rand.Reader, g.bits)
	if err != nil {
		return nil, err
	}

	// private key
	privateData := x509.MarshalPKCS1PrivateKey(privateKey)
	privateBlock := &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: privateData,
	}
	var prv bytes.Buffer
	err = pem.Encode(&prv, privateBlock)
	if err != nil {
		return nil, err
	}

	// public key
	publicKey := &privateKey.PublicKey
	publicData := x509.MarshalPKCS1PublicKey(publicKey)
	publicBlock := &pem.Block{
		Type:  "RSA PUBLIC KEY",
		Bytes: publicData,
	}
	var pub bytes.Buffer
	err = pem.Encode(&pub, publicBlock)
	if err != nil {
		return nil, err
	}

	return &RsaObj{
		Private: prv.Bytes(),
		Public:  pub.Bytes(),
	}, nil
}
