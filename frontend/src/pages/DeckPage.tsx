import { useState } from 'react'
import {
  Box,
  Typography,
  Grid,
  Card,
  CardContent,
  CardActions,
  Button,
  Dialog,
  DialogTitle,
  DialogContent,
  DialogActions,
  TextField,
} from '@mui/material'

interface Deck {
  id: string
  name: string
  cardCount: number
}

export default function DeckPage() {
  const [decks, setDecks] = useState<Deck[]>([])
  const [open, setOpen] = useState(false)
  const [newDeckName, setNewDeckName] = useState('')

  const handleCreateDeck = () => {
    if (newDeckName.trim()) {
      const newDeck: Deck = {
        id: Date.now().toString(),
        name: newDeckName,
        cardCount: 0,
      }
      setDecks([...decks, newDeck])
      setNewDeckName('')
      setOpen(false)
    }
  }

  return (
    <Box>
      <Box sx={{ display: 'flex', justifyContent: 'space-between', mb: 4 }}>
        <Typography variant="h4">My Decks</Typography>
        <Button variant="contained" color="primary" onClick={() => setOpen(true)}>
          Create New Deck
        </Button>
      </Box>

      <Grid container spacing={3}>
        {decks.map((deck) => (
          <Grid item xs={12} sm={6} md={4} key={deck.id}>
            <Card>
              <CardContent>
                <Typography variant="h5" gutterBottom>
                  {deck.name}
                </Typography>
                <Typography color="text.secondary">
                  {deck.cardCount} cards
                </Typography>
              </CardContent>
              <CardActions>
                <Button size="small">Edit</Button>
                <Button size="small" color="error">
                  Delete
                </Button>
              </CardActions>
            </Card>
          </Grid>
        ))}
      </Grid>

      <Dialog open={open} onClose={() => setOpen(false)}>
        <DialogTitle>Create New Deck</DialogTitle>
        <DialogContent>
          <TextField
            autoFocus
            margin="dense"
            label="Deck Name"
            fullWidth
            value={newDeckName}
            onChange={(e) => setNewDeckName(e.target.value)}
          />
        </DialogContent>
        <DialogActions>
          <Button onClick={() => setOpen(false)}>Cancel</Button>
          <Button onClick={handleCreateDeck} variant="contained">
            Create
          </Button>
        </DialogActions>
      </Dialog>
    </Box>
  )
} 