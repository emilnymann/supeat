import { withMaterialColors } from 'tailwind-material-colors';

/** @type {import('tailwindcss').Config} */
const config = {
    content: ['./src/**/*.{html,js,svelte,ts}'],
    theme: {
        extend: {},
    },
    plugins: [],
}

const configWithMaterialColors = withMaterialColors(config, {
    primary: '#94bd39',
});

export default configWithMaterialColors;
