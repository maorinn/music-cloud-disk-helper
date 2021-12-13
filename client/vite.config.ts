import { defineConfig } from 'vite'
import reactRefresh from '@vitejs/plugin-react-refresh'
import tsconfigPaths from 'vite-tsconfig-paths'

// https://vitejs.dev/config/
export default defineConfig({
  plugins: [reactRefresh(), tsconfigPaths()],
  // server: {
  //   proxy: {
  //     //这里是通过请求/api 来转发到 https://api.pingping6.com/
  //     //假如你要请求https://api.*.com/a/a
  //     //那么axios的url，可以配置为 /api/a/a
  //     '/b23': {
  //       target:"https://b23.tv",
  //       changeOrigin:true,
  //       headers:{
  //         "User-Agent":"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.55 Safari/537.36"
  //       },
  //       rewrite: (path) => path.replace(/^\/b23/, ""),
  //     },
  //   }
  // }
})
