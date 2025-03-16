import { render, screen } from '@testing-library/react'
import { describe, it, expect } from 'vitest'
import { ThemeProvider, createTheme } from '@mui/material'
import HomePage from '../HomePage'

const theme = createTheme()

describe('HomePage', () => {
  it('renders Home Page text', () => {
    render(
      <ThemeProvider theme={theme}>
        <HomePage />
      </ThemeProvider>
    )
    expect(screen.getByText('Home Page')).toBeInTheDocument()
  })
}) 