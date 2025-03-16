import React from 'react'
import ReactDOM from 'react-dom/client'
import { createBrowserRouter, RouterProvider } from 'react-router'
import { CssBaseline, ThemeProvider, createTheme } from '@mui/material'
import App from './App'
import ErrorPage from './pages/ErrorPage'
import HomePage from './pages/HomePage'
import LoginPage from './pages/LoginPage'
import MenuPage from './pages/MenuPage'

const router = createBrowserRouter([
  {
    path: "/",
    element: <App />,
    errorElement: <ErrorPage />,
    children: [
      {
        path: "",
        element: <HomePage />,
      },
      {
        path: "login",
        element: <LoginPage />,
      },
      {
        path: "menu",
        element: <MenuPage />,
      },
    ],
  },
])

const theme = createTheme({
  palette: {
    mode: 'dark',
    primary: {
      main: '#ff9800',
    },
    secondary: {
      main: '#4caf50',
    },
  },
})

ReactDOM.createRoot(document.getElementById('root')!).render(
  <React.StrictMode>
    <ThemeProvider theme={theme}>
      <CssBaseline />
      <RouterProvider router={router} />
    </ThemeProvider>
  </React.StrictMode>,
) 