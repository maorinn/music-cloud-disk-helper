import axios from 'axios'
import { config } from 'config/config'
import { action, makeObservable, observable } from 'mobx'
const storage = window.localStorage
class BiliStore {
  @observable qrUrl: string = '' // 二维码url
  @observable oauthKey: string = '' // 鉴权key
  constructor() {
    makeObservable(this)
  }
  // 拉取登录二维码信息
  @action.bound
  async updateLoginQRInfo() {
    let { data } = await axios.request({
      method: 'GET',
      url: `http://${config.SERVER_HOME}/api/v1/bili/LoginUrl`
    })
    data =  data.data
    console.log("设置bili二维码->",data.url);
    
    this.qrUrl = data.url
    this.oauthKey = data.oauthKey
  }
  // 获取登录状态
  @action.bound
  async checkEwmStatus() {
    let { data } = await axios.request({
      method: 'POST',
      url: `http://${config.SERVER_HOME}/api/v1/bili/LoginInfo`,
      data: {
        oauthKey: this.oauthKey
      }
    })
    data =  data.data
    const { status, data: _data } = data
    console.log('哔哩哔哩扫码登录状态->', { status, _data })
    if (status) {
      // 登录成功保存token
      const cookieStr = _data.url.replace('https://passport.biligame.com/crossDomain?', '')
      storage.setItem('bili_cookie_str', cookieStr)
    }
  }
}
export default BiliStore
