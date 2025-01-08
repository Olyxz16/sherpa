import { ref } from 'vue';
import { defineStore } from 'pinia';
import { saveFile, fetchFile } from '@/lib/files'

export const useWorkstationStore = defineStore('workStationStore', () => {
  const currRepoName = ref('');
  const currFileName = ref('');
  const fileContent = ref('');

  function setCurrentRepo(repoName : string) {
    currRepoName.value = repoName;
    currFileName.value = '.env';
    fetchCurrentFile()
    .then(val => {
        fileContent.value = val;
      });
  }
  function setCurrentFile(fileName : string) {
    currFileName.value = fileName;
    fetchCurrentFile()
    .then(val => {
        fileContent.value = val;
      });
  }
  function updateFileContent(payload: string | number) {
    fileContent.value = payload.toString();
  };

  async function saveCurrentFile() {
    // Incorporate source in repo data
    const source = "github.com";
    const repoName = currRepoName.value;
    const fileName = currFileName.value;
    const content = fileContent.value;
    try {
      const res = await saveFile("github.com", repoName, fileName, content);
    } catch(e) {
      // Display error
      console.log("File not saved");
    }
  }
  async function fetchCurrentFile(): Promise<string> {
    const source = "github.com";
    const repoName = currRepoName.value;
    const fileName = currFileName.value;
    const res = await fetchFile(source, repoName, fileName);
    return res;
  }
  return { currRepoName, currFileName, fileContent, setCurrentRepo, setCurrentFile, updateFileContent, saveCurrentFile, fetchCurrentFile };
});
