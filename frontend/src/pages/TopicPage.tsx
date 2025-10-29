import { useEffect, useState } from 'react';
import { useParams } from 'react-router-dom';
import { Box, Container, Typography, Paper, CircularProgress, Alert, Chip } from '@mui/material';
import { financeAPI } from '../services/api';
import type { FinanceTopic } from '../types';

export default function TopicPage() {
  const { id } = useParams<{ id: string }>();
  const [topic, setTopic] = useState<FinanceTopic | null>(null);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState<string | null>(null);

  useEffect(() => {
    const fetchTopic = async () => {
      if (!id) return;
      
      try {
        const data = await financeAPI.getTopicById(id);
        setTopic(data);
      } catch (err) {
        setError(err instanceof Error ? err.message : 'Failed to fetch topic');
      } finally {
        setLoading(false);
      }
    };

    fetchTopic();
  }, [id]);

  if (loading) {
    return (
      <Container maxWidth="lg">
        <Box sx={{ display: 'flex', justifyContent: 'center', py: 8 }}>
          <CircularProgress />
        </Box>
      </Container>
    );
  }

  if (error || !topic) {
    return (
      <Container maxWidth="lg">
        <Box sx={{ my: 4 }}>
          <Alert severity="error">{error || 'Topic not found'}</Alert>
        </Box>
      </Container>
    );
  }

  return (
    <Container maxWidth="lg">
      <Box sx={{ my: 4 }}>
        <Typography variant="h3" component="h1" gutterBottom sx={{ fontWeight: 700 }}>
          {topic.title}
        </Typography>

        <Typography variant="h6" color="text.secondary" paragraph>
          {topic.description}
        </Typography>

        <Paper sx={{ p: 3, mb: 4 }}>
          <Typography variant="h5" gutterBottom sx={{ fontWeight: 600 }}>
            Overview
          </Typography>
          <Typography variant="body1" paragraph>
            {topic.summary}
          </Typography>

          <Box sx={{ mb: 2 }}>
            <Typography variant="subtitle1" fontWeight={600} gutterBottom>
              Key Concepts:
            </Typography>
            <Box sx={{ display: 'flex', gap: 1, flexWrap: 'wrap' }}>
              {topic.keywords.map((keyword) => (
                <Chip key={keyword} label={keyword} />
              ))}
            </Box>
          </Box>
        </Paper>

        <Paper sx={{ p: 3 }}>
          <Typography variant="h5" gutterBottom sx={{ fontWeight: 600 }}>
            Learning Resources
          </Typography>
          <Typography component="div">
            <ul>
              {topic.resources.map((resource, index) => (
                <li key={index}>
                  <a href={resource} target="_blank" rel="noopener noreferrer">
                    {resource}
                  </a>
                </li>
              ))}
            </ul>
          </Typography>
        </Paper>
      </Box>
    </Container>
  );
}
