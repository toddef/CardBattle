import {
  Typography,
  Box,
  Container,
  Grid,
  Card,
  CardContent,
  CardActions,
  Button
} from '@mui/material'

interface MenuOption {
  title: string
  description: string
  action: string
  disabled?: boolean
}

const menuOptions: MenuOption[] = [
  {
    title: 'Quick Match',
    description: 'Start a game against a random opponent',
    action: 'Play Now',
    disabled: true
  },
  {
    title: 'Deck Builder',
    description: 'Create and customize your battle deck',
    action: 'Build Deck',
    disabled: true
  },
  {
    title: 'Profile',
    description: 'View and edit your profile',
    action: 'View Profile',
    disabled: true
  },
  {
    title: 'Leaderboard',
    description: 'See the top players',
    action: 'View Rankings',
    disabled: true
  }
]

export default function MenuPage() {
  return (
    <Container maxWidth="lg">
      <Box sx={{ mt: 4, mb: 8 }}>
        <Typography variant="h3" component="h1" gutterBottom align="center">
          Game Menu
        </Typography>

        <Grid container spacing={4} sx={{ mt: 2 }}>
          {menuOptions.map((option) => (
            <Grid item xs={12} sm={6} md={4} key={option.title}>
              <Card sx={{ height: '100%', display: 'flex', flexDirection: 'column' }}>
                <CardContent sx={{ flexGrow: 1 }}>
                  <Typography variant="h5" component="h2" gutterBottom>
                    {option.title}
                  </Typography>
                  <Typography variant="body2" color="text.secondary">
                    {option.description}
                  </Typography>
                </CardContent>
                <CardActions>
                  <Button 
                    size="large" 
                    fullWidth 
                    variant="contained"
                    disabled={option.disabled}
                  >
                    {option.action}
                  </Button>
                </CardActions>
              </Card>
            </Grid>
          ))}
        </Grid>
      </Box>
    </Container>
  )
} 