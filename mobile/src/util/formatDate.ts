export function formatDateTime(date: string | string) {
  if (!date) return null;
  const dt = new Date(date);
  return Intl.DateTimeFormat('pt-BR', {
    dateStyle: 'short',
    timeStyle: 'short',
  }).format(dt);
}

export function formatDate(date: string | string) {
  if (!date) return null;
  const dt = new Date(date);
  return Intl.DateTimeFormat('pt-BR', {
    dateStyle: 'short',
  }).format(dt);
}
