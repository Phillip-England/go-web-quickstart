
## Cloning the Repository
Run this command to clone the repository. Make sure to replace <project-dir> with the actual path to the directory of your project.
```bash
git clone https://github.com/phillip-england/go-web-quickstart <project-dir>
```

## Installing and Running Tailwind
Run the following command to install tailwind via npm:
```bash
npm install
```
Then, you can run this command to watch for changes in tailwind classnames during development.
```bash
npx tailwindcss -i ./public/input.css -o ./public/output.css --watch
```

