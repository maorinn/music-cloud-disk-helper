import React from 'react'
import { BrowserRouter as Router, Switch, Route } from 'react-router-dom'
import { Home, Login } from '@pages'
import { RouteConfig, renderRoutes } from 'react-router-config'

const routesConfig: RouteConfig[] = [
  {
    path: '/',
    exact: true,
    component: Home
  },
  {
    path: '/login',
    exact: true,
    component: Login
  },
]

interface Props {}

const Routes = (props: Props) => {
  return (
    <Router>
      <Switch>{renderRoutes(routesConfig)}</Switch>
    </Router>
  )
}

export default Routes
