import styled from "styled-components";

const DownloadButton = styled.button`
  height: 80px;
  width: 200px;
  background-color: #d1ce9e;
  font-size: 1em;
  margin: 1em;
  padding: 1em;
  text-transform: uppercase;
  line-height: 30px;
  border: 10px solid #1c6ea4;
  box-shadow: 2px 2px 7px 1px #1c6ea4;
`;

const DisplayPaper = styled.div.attrs({
  children: `
    Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor
    incididunt ut labore et dolore magna aliqua. Ac feugiat sed lectus vestibulum
    mattis ullamcorper velit sed. Maecenas pharetra convallis posuere morbi leo
    urna. Quis varius quam quisque id diam vel. Quisque sagittis purus sit amet
    volutpat. Venenatis a condimentum vitae sapien pellentesque habitant morbi
    tristique. Et odio pellentesque diam volutpat commodo. Semper viverra nam libero
    justo laoreet. Laoreet id donec ultrices tincidunt arcu non sodales neque. Eget
    nulla facilisi etiam dignissim diam quis enim. Aliquam ut porttitor leo a diam
    sollicitudin tempor id eu. A condimentum vitae sapien pellentesque. Massa id
    neque aliquam vestibulum morbi blandit cursus risus. Donec ultrices tincidunt
    arcu non. Posuere sollicitudin aliquam ultrices sagittis orci a scelerisque
    purus. Volutpat lacus laoreet non curabitur gravida.
    
    Id venenatis a condimentum vitae sapien pellentesque habitant. Lectus sit amet
    est placerat in. Diam phasellus vestibulum lorem sed risus ultricies tristique
    nulla aliquet. Morbi leo urna molestie at elementum eu facilisis sed. Dolor
    purus non enim praesent. Sollicitudin tempor id eu nisl nunc. Est velit egestas
    dui id ornare arcu odio. Id cursus metus aliquam eleifend mi in. Velit egestas
    dui id ornare. Massa sed elementum tempus egestas sed sed risus. Lacus sed
    viverra tellus in hac habitasse platea dictumst vestibulum. Risus commodo
    viverra maecenas accumsan. Amet nisl suscipit adipiscing bibendum est. Eget mi
    proin sed libero enim sed faucibus. Dolor morbi non arcu risus quis. Bibendum ut
    tristique et egestas quis ipsum suspendisse ultrices gravida.
    
    Nulla pellentesque dignissim enim sit amet venenatis urna. Vitae proin sagittis
    nisl rhoncus mattis. Nisi lacus sed viverra tellus in hac habitasse platea.
    Fusce ut placerat orci nulla pellentesque. Nec feugiat in fermentum posuere urna
    nec tincidunt praesent semper. Dignissim cras tincidunt lobortis feugiat vivamus
    at. Laoreet non curabitur gravida arcu ac tortor dignissim. Aliquet lectus proin
    nibh nisl condimentum id venenatis a. Odio aenean sed adipiscing diam donec.
    Ipsum a arcu cursus vitae congue mauris. Molestie ac feugiat sed lectus.
    Ullamcorper malesuada proin libero nunc consequat interdum varius sit.
    Ullamcorper velit sed ullamcorper morbi tincidunt ornare. A diam maecenas sed
    enim. Nisl suscipit adipiscing bibendum est ultricies integer quis. Hendrerit
    dolor magna eget est. Feugiat pretium nibh ipsum consequat nisl. Fermentum
    posuere urna nec tincidunt praesent semper feugiat nibh.
    
    Ipsum dolor sit amet consectetur adipiscing elit pellentesque habitant. Tortor
    pretium viverra suspendisse potenti nullam ac tortor. Nunc pulvinar sapien et
    ligula ullamcorper. Venenatis lectus magna fringilla urna porttitor rhoncus.
    Massa eget egestas purus viverra accumsan in. Sit amet porttitor eget dolor. In
    metus vulputate eu scelerisque felis imperdiet proin fermentum leo. In hac
    habitasse platea dictumst vestibulum rhoncus est pellentesque. Pellentesque sit
    amet porttitor eget dolor morbi non arcu risus. Sit amet nisl suscipit
    adipiscing bibendum. Rutrum quisque non tellus orci ac. Nullam vehicula ipsum a
    arcu cursus vitae congue mauris. Leo duis ut diam quam nulla porttitor. Feugiat
    pretium nibh ipsum consequat.
    
    Egestas dui id ornare arcu odio. Tristique senectus et netus et malesuada fames
    ac turpis egestas. Tincidunt augue interdum velit euismod in pellentesque massa
    placerat. Proin sagittis nisl rhoncus mattis rhoncus urna neque. Iaculis at erat
    pellentesque adipiscing commodo elit at. Cras fermentum odio eu feugiat pretium
    nibh ipsum consequat nisl. Nec sagittis aliquam malesuada bibendum arcu vitae.
    Cursus vitae congue mauris rhoncus. Tellus at urna condimentum mattis
    pellentesque id nibh. Porta nibh venenatis cras sed felis eget velit aliquet.
    Condimentum id venenatis a condimentum vitae sapien pellentesque habitant.
  `
})`
  height: 100%;
  width: 100%;
  padding: 20px;
  margin: 20px;
  background-color: #dac3c3;
  white-space: pre-wrap;
`;

const Home = () => {
  return (
    <div>
      <DownloadButton>Download</DownloadButton>
      <div style={{ height: "30px" }} />
      <DisplayPaper />
    </div>
  );
};

export default Home;
