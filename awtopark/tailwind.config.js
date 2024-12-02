/** @type {import('tailwindcss').Config} */
export default {
  content: ['./src/**/*.{html,js,svelte,ts}'],
  theme: {
    extend: {},
  },
  plugins: [
    require('tailwind-scrollbar'), // Scrollbar eklentisini ekliyoruz
  ],
  variants: {
    scrollbar: ['rounded'], // Yuvarlatılmış köşeler için
  },
}
