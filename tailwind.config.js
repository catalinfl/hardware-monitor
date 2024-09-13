module.exports = {
  mode: 'jit',
  content: [
    './**/*.{html,js,jsx,ts,tsx,templ}', // Adjust the paths according to your project structure
    './**/*.{templ}', // Include other paths as needed
  ],
  theme: {
    extend: {},
  },
  plugins: [],
}