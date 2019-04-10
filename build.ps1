$env:GOOS = "linux"

echo "Cleaning"
rm ./backend/backend -Confirm:$false
rm ./backend/public -Confirm:$false
mkdir ./backend/public

echo "Building App"
cd backend
go build
cd ..

echo "Building UI"
cd ui
npm build
cd ..

echo "Copy ui to backend"
cp -r ./ui/build/* ./backend/public

echo "Building Image"
cd backend
docker build -t patnaikshekhar/kubedashboard:0.0.1 .
cd ..
