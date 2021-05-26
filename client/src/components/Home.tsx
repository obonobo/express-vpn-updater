import styled from "styled-components";

const DownloadButton = styled.button`
height:50px;
width:100px;
background-color:red;
margin : 20px;
padding : 10px;
text-transform: uppercase;

`;

const Home = () => {
    return (
      <DownloadButton>Download</DownloadButton>
    );
  };

export default Home ; 