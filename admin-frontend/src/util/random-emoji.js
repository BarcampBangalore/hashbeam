const emojis = [
  { name: 'tractor', text: '🚜' },
  { name: 'frog', text: '🐸' },
  { name: 'caterpillar', text: '🐛' },
  { name: 'crab', text: '🦀' },
  { name: 'watermelon', text: '🍉' }
];

export default function randomEmoji() {
  return emojis[Math.floor(Math.random() * emojis.length)];
}
