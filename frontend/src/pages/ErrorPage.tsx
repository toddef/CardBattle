import { Box, Typography, Button } from '@mui/material'
import { useRouteError, useNavigate } from 'react-router-dom'

export default function ErrorPage() {
  const error = useRouteError() as { statusText?: string; message?: string }
  const navigate = useNavigate()

  return (
    <Box
      sx={{
        display: 'flex',
        flexDirection: 'column',
        alignItems: 'center',
        justifyContent: 'center',
        height: '100vh',
        textAlign: 'center',
      }}
    >
      <Typography variant="h1" color="error" gutterBottom>
        Oops!
      </Typography>
      <Typography variant="h4" gutterBottom>
        Sorry, an unexpected error has occurred.
      </Typography>
      <Typography variant="body1" color="text.secondary" gutterBottom>
        {error.statusText || error.message || 'Unknown error'}
      </Typography>
      <Button
        variant="contained"
        color="primary"
        onClick={() => navigate('/')}
        sx={{ mt: 2 }}
      >
        Go Home
      </Button>
    </Box>
  )
} 