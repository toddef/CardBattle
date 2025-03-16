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
    <Container>
      <Box sx={{ mt: 4 }}>
        <Typography variant="h3">Menu Page</Typography>
      </Box>
    </Container>
  )
} 