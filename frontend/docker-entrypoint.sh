#!/bin/sh

echo "${VITE_API_URL}"

FILE="/usr/share/nginx/html/config/env.js"

if [ ! -f "$FILE" ]; then
  echo "Missing $FILE, generating default config..."

  # Write runtime env config
  echo "window.env = { VITE_API_URL: '${VITE_API_URL}' };" > "/app/config/env.js"

  mv /app/* /usr/share/nginx/html
else 
  echo "window.env = { VITE_API_URL: '${VITE_API_URL}' };" > "/usr/share/nginx/html/config/env.js"
fi

# Start nginx
exec nginx -g "daemon off;"
