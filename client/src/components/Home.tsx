import styled from "styled-components";

const DownloadButton = styled.button.attrs({ children: "Download" })`
  height: 80px;
  width: 200px;
  background-color: #ad0a0a;
  color: white;
  font-size: 1em;
  margin: 1em;
  padding: 1em;
  text-transform: uppercase;
  line-height: 30px;
  border: 10px solid #f1f2f3;
  box-shadow: 2px 2px 7px 1px #a09d9d;
  font-family: Arial, Helvetica, sans-serif;
  font-weight: bold;
  border-radius: 12px;
  transition-duration: 0.4s;
  &:hover {
    background-color: #e20636;
  }
  &:active {
    transform: translateY(4px);
    background-color: #ff0000;
  }
`;

const DisplayPaper = styled.div.attrs({
  children: `
  # Changelog

  ## [1.0.0] - 2017-06-20
  ### Added
  - New visual identity by [@tylerfortune8](https://github.com/tylerfortune8).
  - Version navigation.
  - Links to latest released version in previous versions.
  - "Why keep a changelog?" section.
  - "Who needs a changelog?" section.
  - "How do I make a changelog?" section.
  - "Frequently Asked Questions" section.
  - New "Guiding Principles" sub-section to "How do I make a changelog?".
  - Simplified and Traditional Chinese translations from [@tianshuo](https://github.com/tianshuo).
  - German translation from [@mpbzh](https://github.com/mpbzh) & [@Art4](https://github.com/Art4).
  - Italian translation from [@azkidenz](https://github.com/azkidenz).
  - Swedish translation from [@magol](https://github.com/magol).
  - Turkish translation from [@karalamalar](https://github.com/karalamalar).
  - French translation from [@zapashcanon](https://github.com/zapashcanon).
  - Brazilian Portugese translation from [@Webysther](https://github.com/Webysther).
  - Polish translation from [@amielucha](https://github.com/amielucha) & [@m-aciek](https://github.com/m-aciek).
  - Russian translation from [@aishek](https://github.com/aishek).
  - Czech translation from [@h4vry](https://github.com/h4vry).
  - Slovak translation from [@jkostolansky](https://github.com/jkostolansky).
  - Korean translation from [@pierceh89](https://github.com/pierceh89).
  - Croatian translation from [@porx](https://github.com/porx).
  - Persian translation from [@Hameds](https://github.com/Hameds).
  - Ukrainian translation from [@osadchyi-s](https://github.com/osadchyi-s).
  
  ### Changed
  - Start using "changelog" over "change log" since it's the common usage.
  - Start versioning based on the current English version at 0.3.0 to help
  translation authors keep things up-to-date.
  - Rewrite "What makes unicorns cry?" section.
  - Rewrite "Ignoring Deprecations" sub-section to clarify the ideal
    scenario.
  - Improve "Commit log diffs" sub-section to further argument against
    them.
  - Merge "Why can’t people just use a git log diff?" with "Commit log
    diffs"
  - Fix typos in Simplified Chinese and Traditional Chinese translations.
  - Fix typos in Brazilian Portuguese translation.
  - Fix typos in Turkish translation.
  - Fix typos in Czech translation.
  - Fix typos in Swedish translation.
  - Improve phrasing in French translation.
  - Fix phrasing and spelling in German translation.
  
  ### Removed
  - Section about "changelog" vs "CHANGELOG".`,
})`
  width: 100%;
  padding: 20px;
  margin: 20px;
  background-color: #dac3c3;
  white-space: pre-wrap;
  font-size: 80%;
  text-align: left;
`;

const VersionText = styled.text.attrs({
  children: "Express VPN Latest Version : Version 4.2.0.6.9",
})``;

const Grid = styled.div`
  display: grid;
  grid-template-columns: repeat(2, 1fr);
  justify-content: center;
  align-items: center;
  margin: 3%;
`;

const GridItem = styled.div`
  & > ${DisplayPaper} {
    float: right;
  }
  & > ${DownloadButton} {
    margin-bottom: calc(1em + 30px);
    margin-right: 30%;
  }
`;

const Root = styled.div`
  text-align: center;
  justify-content: center;
  align-items: center;
`;

const Home = () => {
  return (
    <Root>
      <VersionText />
      <Grid>
        <GridItem>
          <DownloadButton />
        </GridItem>
        <GridItem>
          <DisplayPaper />
        </GridItem>
      </Grid>
    </Root>
  );
};

export default Home;
