import {
  Route, 
  useNavigate,
  Routes,
} from 'react-router-dom';
import LoginForm from './components/LoginForm';
import RegisterForm from './components/RegisterForm';
import ImageForm from './components/ImageForm';

function App() {
  const history = useNavigate();

  return (
    <Routes>
      <Route exact path="/" element={LoginForm()} />
      <Route path="/register" element={RegisterForm()} /> 
      <Route path="/upload" element={ImageForm()} />
    </Routes>
  );
}

export default App;
