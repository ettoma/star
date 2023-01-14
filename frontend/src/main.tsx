import React from 'react'
import ReactDOM from 'react-dom/client'
import { RouterProvider } from 'react-router'
import { createBrowserRouter } from 'react-router-dom'
import ErrorPage from './errorPage'
import SignIn from '../pages/signIn'
import Register from '../pages/register'
import Users from '../pages/users'
import Home from '../pages/home'
import { Grommet } from 'grommet'
import { theme } from './globals/styles/grommet'
import Root from './Root'



const router = createBrowserRouter([
  {
    path: "/",
    element: <Root />,
    errorElement: <ErrorPage />,
    children: [
      {
        path: '/',
        element: <Home />
      },
      {
        path: '/signin',
        element: <SignIn />
      },
      {
        path: '/register',
        element: <Register />
      },
      {
        path: "/users",
        element: <Users />
      }

    ]
  },
]);


ReactDOM.createRoot(document.getElementById('root') as HTMLElement).render(
  <Grommet full theme={theme} background="#212121">
    <React.StrictMode>
      <RouterProvider router={router} />
    </React.StrictMode>
  </Grommet>
)
