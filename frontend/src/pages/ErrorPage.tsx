import { useRouteError } from 'react-router-dom'
import { Typography, Box, Container, Button } from '@mui/material'
import { Link as RouterLink } from 'react-router-dom'

interface RouterError {
  statusText?: string
  message?: string
}

export default function ErrorPage() {
  const error = useRouteError() as RouterError

  return (
    <Container maxWidth="sm">
      <Box
        sx={{
          mt: 8,
          display: 'flex',
          flexDirection: 'column',
          alignItems: 'center',
          textAlign: 'center',
        }}
      >
        <Typography variant="h1" component="h1" gutterBottom>
          Oops!
        </Typography>
        <Typography variant="h5" component="h2" gutterBottom color="text.secondary">
          Sorry, an unexpected error has occurred.
        </Typography>
        <Typography variant="body1" color="text.secondary" sx={{ mb: 4 }}>
          {error.statusText || error.message || 'Unknown error'}
        </Typography>
        <Button
          component={RouterLink}
          to="/"
          variant="contained"
          color="primary"
          size="large"
        >
          Return to Home
        </Button>
      </Box>
    </Container>
  )
} 