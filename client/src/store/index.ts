import BiliStore from './biliStore'
import NeteaseStore from './neteaseStore'
let biliStore = new BiliStore()
let neteaseStore = new NeteaseStore()
const stores = {
  biliStore,
  neteaseStore
}
export default stores
