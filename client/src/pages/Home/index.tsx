import { useLocalStore, useObserver } from 'mobx-react'
import React, { useEffect } from 'react'
import stores from '../../store'
import styled from 'styled-components'
import QRCode from 'qrcode.react'
import TextField from '@mui/material/TextField'
import Navigation from '@components/Navigation'
import { Container } from '@components/Container'
import { RowContainer } from '@components/RowContainer'
import { ColumnContainer } from '@components/ColumnContainer'
import Button from '@mui/material/Button'
import Bottom from '@components/Bottom'
import Modal from '@mui/material/Modal'
import Snackbar from '@mui/material/Snackbar'
import { Backdrop, CircularProgress } from '@mui/material'
interface Props {}
const Home = (props: Props) => {
  const biliStore = useLocalStore(() => stores.biliStore)
  const neteaseStore = useLocalStore(() => stores.neteaseStore)
  const [open, setOpen] = React.useState(false)
  const [backdropOpen, setBackdropOpen] = React.useState(false)
  const [snackbarOpen, setSnackbarOpen] = React.useState(false)
  const [snackbarMessage, setSnackbarMessage] = React.useState('')
  const [bvId, setBvId] = React.useState('')
  const handleOpen = () => setOpen(true)
  const handleClose = () => {
    setOpen(false)
    clearInterval(loginCheckTimer)
  }
  let loginCheckTimer: any = null
  const handleBinNetease = async () => {
    await neteaseStore.updateLoginQRInfo()
    handleOpen()
    // 检查 登录状态定时任务
    loginCheckTimer = setInterval(async () => {
      const bool = await neteaseStore.checkEwmStatus()
      if (bool) {
        setSnackbarOpen(true)
        setSnackbarMessage('登录成功')
        handleClose()
      }
    }, 2000)
  }
  const handleUploadSong = async () => {
    setBackdropOpen(true)
    const resp = await neteaseStore.uploadBiliSong(bvId)
    setBackdropOpen(false)

    setSnackbarOpen(true)
    if (resp.status == 200 && resp.data.code == 0) {
      setSnackbarMessage('上传成功')
    } else {
      setSnackbarMessage('上传失败，请稍后再试~')
    }
  }
  return useObserver(() => (
    <Container>
      <Navigation />
      <Snackbar
        anchorOrigin={{ vertical: 'bottom', horizontal: 'center' }}
        open={snackbarOpen}
        onClose={() => {
          setSnackbarOpen(false)
        }}
        autoHideDuration={3000}
        message={snackbarMessage}
        key="1"
      />
      <Backdrop
        sx={{ color: '#fff', zIndex: (theme) => theme.zIndex.drawer + 1 }}
        open={backdropOpen}
        onClick={() => {
          setBackdropOpen(false)
        }}
      >
        <CircularProgress color="inherit" />
      </Backdrop>
      <MainBody>
        <Card>
          <ColumnContainer>
            <TextField
              id="outlined-basic"
              label="BV号"
              variant="standard"
              onChange={(event: React.ChangeEvent<HTMLInputElement>) => {
                setBvId(event.target.value)
              }}
            />
            <Buttons>
              <Button
                variant="contained"
                style={{
                  marginRight: '30px',
                  width: '200px'
                }}
                onClick={handleUploadSong}
              >
                上传云音乐
              </Button>
              <Button
                onClick={handleBinNetease}
                variant="contained"
                style={{
                  width: '50px'
                }}
              >
                {neteaseStore.cookieStr ? '重新登录' : '扫码登录'}
              </Button>
            </Buttons>
          </ColumnContainer>
        </Card>
        <Modal
          open={open}
          onClose={handleClose}
          aria-labelledby="modal-modal-title"
          aria-describedby="modal-modal-description"
        >
          <ModalContent>
            {neteaseStore.qrUrl && <QRCode value={neteaseStore.qrUrl}></QRCode>}
            <div
              style={{
                fontSize: '14px',
                marginTop: '5px'
              }}
            >
              使用网易云音乐App扫码绑定
            </div>
          </ModalContent>
        </Modal>
      </MainBody>
      {/* <Bottom /> */}
    </Container>
  ))
}

const MainBody = styled.div({
  height: '78vh',
  display: 'flex',
  width: '100%',
  flexDirection: 'column'
})
const Card = styled.div({
  display: 'flex',
  width: '100%',
  background: 'rgb(255,255,255,.5)',
  backdropFilter: 'blur(16px) saturate(180%)',
  border: '4px solid #fff',
  borderRadius: '4px solid #fff',
  padding: '30px',
  WebkitBoxShadow: '0 0 15px 0 rgb(0 0 0 / 10%)',
  margin: 'auto',
  position: 'relative',
  top: '35%',
  transform: 'translateY(-50%)',
  transition: '.5s',
  flexDirection: 'column',
  justifyContent: 'center'
})
const Buttons = styled.div({
  display: 'flex',
  flexDirection: 'row',
  marginTop: '20px',
  marginLeft: '20px'
})
const ModalContent = styled.div({
  background: '#fff',
  margin: 'auto',
  textAlign: 'center',
  alignItems: 'center',
  padding: '20px',
  borderRadius: '5px',
  width: '80%',
  marginTop: '60%',
  display: 'flex',
  justifyContent: 'center',
  flexDirection: 'column'
})
export default Home
