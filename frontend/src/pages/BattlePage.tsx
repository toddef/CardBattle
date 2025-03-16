import { Box, Typography, Button, Paper, Grid } from '@mui/material'

export default function BattlePage() {
  return (
    <Box>
      <Typography variant="h4" gutterBottom>
        Battle Arena
      </Typography>

      <Grid container spacing={3}>
        <Grid item xs={12} md={8}>
          <Paper sx={{ p: 2, minHeight: '400px' }}>
            <Typography variant="h6" gutterBottom>
              Game Board
            </Typography>
            <Box sx={{ height: '300px', display: 'flex', alignItems: 'center', justifyContent: 'center' }}>
              <Typography color="text.secondary">
                No active battle. Start a new game or join an existing one.
              </Typography>
            </Box>
          </Paper>
        </Grid>

        <Grid item xs={12} md={4}>
          <Paper sx={{ p: 2 }}>
            <Typography variant="h6" gutterBottom>
              Battle Options
            </Typography>
            <Box sx={{ display: 'flex', flexDirection: 'column', gap: 2 }}>
              <Button variant="contained" color="primary" fullWidth>
                Quick Match
              </Button>
              <Button variant="outlined" color="primary" fullWidth>
                Create Private Game
              </Button>
              <Button variant="outlined" color="primary" fullWidth>
                Join Private Game
              </Button>
            </Box>
          </Paper>

          <Paper sx={{ p: 2, mt: 2 }}>
            <Typography variant="h6" gutterBottom>
              Active Games
            </Typography>
            <Typography color="text.secondary">
              No active games found
            </Typography>
          </Paper>
        </Grid>
      </Grid>
    </Box>
  )
} 