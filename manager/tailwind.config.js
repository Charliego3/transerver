import forms from "@tailwindcss/forms";

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
                ibm: "'IBMPlexMono', monospace",
                splash: "'Splash', cursive",
            },
            textColor: {
                skin: {
                    base: withOpacity('--text-color'),
                    accent: withOpacity('--accent-color'),
                    second: withOpacity('--second-color'),
                }
            },
            backgroundColor: {
                skin: {
                    fill: withOpacity('--color-background'),
                    accent: withOpacity('--accent-color'),
                    second: withOpacity('--second-color'),
                }
            },
            borderColor: {
                skin: {
                    fill: withOpacity('--color-background'),
                    accent: withOpacity('--accent-color'),
                }
            },
            backgroundImage: {
                'image-dark': "url('images/background-dark.jpg')",
                'image-light': "url('images/background.jpg')",
            },
            gradientColorStops: {
                skin: {
                    background: withOpacity('--color-background'),
                }
            },
            ringColor: {
                skin: {
                    accent: withOpacity('--accent-color'),
                }
            },
            fill: {
                skin: {
                    fill: withOpacity('--accent-color')
                }
            }
        },
    },
    plugins: [forms],
}
