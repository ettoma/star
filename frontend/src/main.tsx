import React from 'react'
import ReactDOM from 'react-dom/client'
import { RouterProvider } from 'react-router'
import { createBrowserRouter } from 'react-router-dom'
import SignIn from '../pages/signIn'
import Register from '../pages/register'
import Users from '../pages/users'
import './index.css'
import Home from '../pages/home'


const router = createBrowserRouter([
  {
    path:"/",
    element: <Home />,
    errorElement: <h2>Page not found</h2>
  },
  {
    path: "/signin",
    element: <SignIn />,
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
  <main className='main_container'>
    <div id="star_title">
      <h2>Star</h2>
    </div>
    <React.StrictMode>
      <RouterProvider router={router} />
    </React.StrictMode>
  </main>
)
