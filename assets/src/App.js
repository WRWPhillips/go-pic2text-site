import logo from './logo.svg';
import './App.css';

function App() {
  return (
    <div className="App">
    <h1>Welcome to go-pic2text!</h1>
    <h3>Enter image and params here</h3>
    <form
      enctype="multipart/form-data"
      action="http://localhost:9000/api/upload"
      method="post"
    >
      <input type="file" name="myFile" />
      <input type="number" name="width" value={80}/>
      <input type="number" name="height" value={25}/>
      <input type="string" name="palette" value="'$@B%8&WM#*oahkbdpqwmZO0QLCJUYXzcvunxrjft/\\|()1{}[]?-_+~<>i!lI;:,\^`."/>
      <input type="submit" value="upload" />
    </form>
    <h3>Enter login details here </h3>
    <form 
      enctype="application/x-www-form-urlencoded"
      action="http://localhost:9000/api/signin"
      method="post"
    >
      <label>
        Name:
      <input type="string" name="Username" />
      </label>
      <label>
        Password:
      <input type="password" name="Password" />
      </label>
      <input type="submit" value="submit" />
    </form>
    <h3>Sign up here </h3>
    <form 
      enctype="application/x-www-form-urlencoded"
      action="http://localhost:9000/api/signup"
      method="post"
    >
      <label>
        Name:
      <input type="string" name="Username" />
      </label>
      <label>
        Password:
      <input type="password" name="Password" />
      </label>
      <input type="submit" value="submit" />
    </form>
    </div>
  );
}

export default App;
