import React from 'react'
import { FunctionComponent } from 'react'
import styled from 'styled-components'
interface Props {}

const Navigation: FunctionComponent<Props> = () => {
  return (
    <Head>
      <LogoText>Netease Cloud Kit</LogoText>
    </Head>
  )
}
const Head = styled.div({
  display: 'flex',
  flexDirection: 'column',
  width: '85vw',
  margin: 'auto',
  paddingTop: '10px'
})
const LogoText = styled.div({
  width: '100%',
  paddingLeft: 0,
  margin: 'auto',
  fontSize: '28px',
  paddingTop: '10px',
  paddingBottom: '10px',
  fontWeight: 500,
  textAlign: 'center'
})

export default Navigation
