/** @type {import('tailwindcss').Config} */
module.exports = {
  content: [
    "./App.{js,jsx,ts,tsx}",
    "./app/**/*.{js,jsx,ts,tsx}",
    "./components/**/*.{js,jsx,ts,tsx}"
  ],
  presets: [require("nativewind/preset")],
  theme: {
    extend: {
      colors: {
        primary: {
          DEFAULT: '#0060FF',
          50: '#E6F0FF',
          100: '#CCE0FF',
          500: '#0060FF',
          600: '#0050D6',
          700: '#0040AD',
        },
      },
    },
  },
  plugins: [],
}
