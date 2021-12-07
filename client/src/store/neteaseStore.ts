import axios from 'axios'
import { action, makeObservable, observable } from 'mobx'
const storage = window.localStorage
class NeteaseStore {
  constructor() {
    makeObservable(this)
  }
  @observable qrUrl: string = '' // 二维码url 二维码url 规则：https://music.163.com/login?codekey=${query.key}
  @observable key: string = 'null' // 鉴权key
  // 拉取登录二维码信息
  @action.bound
  async updateLoginQRInfo() {
    let { data } = await axios.request({
      method: 'GET',
      url: `http://127.0.0.1:8000/api/v1/netease/qrKey`
    })
    data =  data.data
    this.qrUrl = `https://music.163.com/login?codekey=${data.key}`
    this.key = data.key
  }
  // 获取登录状态
  @action.bound
  async checkEwmStatus() {
    let { data } = await axios.request({
      method: 'GET',
      url: `http://127.0.0.1:8000/api/v1/netease/checkQr`,
      params: {
        key: this.key
      }
    })
    data =  data.data
    const { status, cookie } = data
    console.log('网易云扫码登录状态->', { status, cookie })
    if (status) {
      // 登录成功保存token
      storage.setItem('netease_cookie_str', cookie)
    }
  }
}
export default NeteaseStore