const fs = require('fs');

try {
  const data = fs.readFileSync('./users.json');
  const users = JSON.parse(data);
  users.forEach(u => console.log(u.name));
} catch (err) {
  console.error('Error reading file', err);
}