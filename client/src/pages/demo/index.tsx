import { useLocalStore, useObserver } from 'mobx-react'
import React from 'react'
import stores from '../../store'
import styled from 'styled-components'
import QRCode from 'qrcode.react'
import { Container } from '@components/Container'
interface Props {}
const Demo = (props: Props) => {
  const biliStore = useLocalStore(() => stores.biliStore)
  const neteaseStore = useLocalStore(() => stores.neteaseStore)
  return useObserver(() => (
    <Container>
      <QR>
        哔哩哔哩
        {biliStore.qrUrl&&<QRCode value={biliStore.qrUrl}></QRCode>}
        
        <Button onClick={biliStore.updateLoginQRInfo}>获取二维码</Button>
        <Button onClick={biliStore.checkEwmStatus}>检查二维码</Button>
      </QR>
      <QR>
        网易云音乐
        {neteaseStore.qrUrl&&<QRCode value={neteaseStore.qrUrl}></QRCode>}
        
        <Button onClick={neteaseStore.updateLoginQRInfo}>获取二维码</Button>
        <Button onClick={neteaseStore.checkEwmStatus}>检查二维码</Button>
      </QR>
    </Container>
  ))
}

export default Demo



const QR = styled.div({
  display: 'flex',
  flexGrow: 1,
  width: '600px',
  height: '600px',
  flexDirection: 'row'
})

const Button = styled.button({
  color: '#cccc00',
  backgroundColor: '#ccccc',
  margin: 10
})
