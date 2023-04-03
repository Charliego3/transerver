function withOpacity(variableName) {
    return ({opacityValue}) => {
        if (opacityValue !== undefined) {
            return `rgba(var(${variableName}), ${opacityValue})`
        }
        return `rgb(var(${variableName}))`
    }
}

/** @type {import('tailwindcss').Config} */
export default {
    content: ['./src/**/*.{html,js,svelte,ts}'],
    theme: {
        extend: {
            fontFamily: {
                dm: "'DM Mono'",
                splash: "'Splash', cursive",
            },
            textColor: {

            },
            backgroundColor: {
                skin: {
                    fill: withOpacity('--color-background'),
                }
            },
            backgroundImage: {
                'image-dark': "url('/images/background-dark.jpg')",
                'image-light': "url('/images/background.jpg')",
            },
            gradientColorStops: {
                skin: {
                    background: withOpacity('--color-background'),
                }
            },
        },
    },
    plugins: [],
}
