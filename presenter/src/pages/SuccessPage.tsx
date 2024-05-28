import { useParams } from "react-router-dom";

function SuccessPage() {
  const { subject } = useParams() ;
  return (
    <div className="flex flex-col justify-center align-middle"> 
      <h1>Success</h1>
      <p>Your {subject} was successful!</p>
    </div>
  );
}

export default SuccessPage;