import { Outlet, Link as RouterLink } from 'react-router-dom'
import {
  AppBar,
  Toolbar,
  Typography,
  Button,
  Box,
  Container,
  Stack
} from '@mui/material'

export default function App() {
  return (
    <Box sx={{ display: 'flex', flexDirection: 'column', minHeight: '100vh' }}>
      <AppBar position="static" elevation={0}>
        <Toolbar>
          <Typography variant="h6" component={RouterLink} to="/" sx={{ 
            flexGrow: 1, 
            textDecoration: 'none',
            color: 'inherit'
          }}>
            Card Battle
          </Typography>
          <Stack direction="row" spacing={2}>
            <Button color="inherit" component={RouterLink} to="/">
              Home
            </Button>
            <Button color="inherit" component={RouterLink} to="/login">
              Login
            </Button>
            <Button color="inherit" component={RouterLink} to="/menu">
              Menu
            </Button>
          </Stack>
        </Toolbar>
      </AppBar>

      <Container component="main" sx={{ mt: 4, mb: 4, flex: 1 }}>
        <Outlet />
      </Container>

      <Box component="footer" sx={{ py: 3, px: 2, mt: 'auto', backgroundColor: 'background.paper' }}>
        <Container maxWidth="sm">
          <Typography variant="body2" color="text.secondary" align="center">
            © {new Date().getFullYear()} Card Battle Game
          </Typography>
        </Container>
      </Box>
    </Box>
  )
} 