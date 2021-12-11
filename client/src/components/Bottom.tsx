import React from 'react'
import { FunctionComponent } from 'react'
import styled from 'styled-components'
interface Props {}

const Bottom: FunctionComponent<Props> = () => {
  return (
    <Container>
      <Text>2021</Text>
    </Container>
  )
}
const Container = styled.div({
  display: 'flex',
  flexDirection: 'column',
  width: '100%',
  margin: 'auto',
  paddingTop: '10px',
  backgroundColor:"#000000"
})
const Text = styled.div({
  width: '100%',
  paddingLeft: 0,
  margin: 'auto',
  fontSize: '14px',
  color: "#fff",
  paddingTop: '5px',
  paddingBottom: '5px',
  fontWeight: 500,
  textAlign: 'center'
})

export default Bottom
