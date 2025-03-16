import { render, screen } from '@testing-library/react'
import { describe, it, expect } from 'vitest'
import { ThemeProvider } from '@mui/material'
import { theme } from '../../theme'
import HomePage from '../HomePage'

const TestWrapper = ({ children }: { children: React.ReactNode }) => (
  <ThemeProvider theme={theme}>
    {children}
  </ThemeProvider>
)

describe('HomePage', () => {
  const renderWithTheme = (ui: React.ReactElement) => {
    return render(ui, {
      wrapper: TestWrapper
    })
  }

  it('renders welcome message', () => {
    renderWithTheme(<HomePage />)

    expect(screen.getByText('Welcome to Card Battle')).toBeInTheDocument()
    expect(screen.getByText('Begin Your Adventure')).toBeInTheDocument()
  })

  it('renders game description', () => {
    renderWithTheme(<HomePage />)

    expect(screen.getByText(/Card Battle is an exciting strategic card game/)).toBeInTheDocument()
    expect(screen.getByText(/Login or create an account/)).toBeInTheDocument()
  })
}) 