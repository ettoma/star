import React from 'react'
import ReactDOM from 'react-dom/client'
import { RouterProvider } from 'react-router'
import { createBrowserRouter } from 'react-router-dom'
import SignIn from '../pages/signIn'
import Register from '../pages/register'
import Users from '../pages/users'
import './index.css'


const router = createBrowserRouter([
  {
    path: "/",
    element: <SignIn />,
    errorElement: <h2>Page not found</h2>
  },
  {
    path: "/register",
    element: <Register />
  },
  {
    path: "/users",
    element: <Users />
  }
]);

ReactDOM.createRoot(document.getElementById('root') as HTMLElement).render(
  <React.StrictMode>
    <RouterProvider router={router} />
  </React.StrictMode>
)
