import { Typography, Box, Paper, Container } from '@mui/material'

export default function HomePage() {
  return (
    <Container maxWidth="md">
      <Box sx={{ mt: 4, mb: 8 }}>
        <Typography variant="h1" component="h1" gutterBottom align="center">
          Welcome to Card Battle
        </Typography>
        <Typography variant="h2" component="h2" gutterBottom align="center" color="primary">
          Begin Your Adventure
        </Typography>
        <Paper elevation={3} sx={{ p: 4, mt: 4 }}>
          <Typography variant="body1" paragraph>
            Card Battle is an exciting strategic card game where you can build your deck,
            battle other players, and become the ultimate champion.
          </Typography>
          <Typography variant="body1">
            Login or create an account to start your journey!
          </Typography>
        </Paper>
      </Box>
    </Container>
  )
} 