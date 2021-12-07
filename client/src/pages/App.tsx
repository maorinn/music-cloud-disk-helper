import { Routes } from '@config'
import React, { useState } from 'react'
import './App.css'
import { Provider } from 'mobx-react'
import stores from '../store'
function App() {
  return (
    <Provider {...stores}>
      <Routes />
    </Provider>
  )
}

export default App
