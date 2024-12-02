import { defineConfig } from 'vite'
import { svelte } from '@sveltejs/vite-plugin-svelte'

// https://vite.dev/config/
export default defineConfig({
  plugins: [svelte()],
  server: {
    host: true, // Tüm ağ arayüzlerinden erişimi etkinleştirir
    port: 5173, // Port numarasını belirleyin (isteğe bağlı)
    strictPort: true, // Port zaten doluysa hata verir
  },
})
