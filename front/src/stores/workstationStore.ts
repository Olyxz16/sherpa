import { ref } from 'vue';
import { defineStore } from 'pinia';
import { saveFile, fetchFile } from '@/lib/files'

export const useWorkstationStore = defineStore('workStationStore', () => {
  const currRepoNameRef = ref('');
  const currFileNameRef = ref('');
  const fileContentRef = ref('');

  const repoName = () => currRepoNameRef.value;
  const fileName = () => currFileNameRef.value;
  const fileContent = () => fileContentRef.value;

  function setCurrentRepo(repoName : string) {
    currRepoNameRef.value = repoName;
    currFileNameRef.value = '.env';
    fileContentRef.value = ''; // Fetch file data here ?
  }
  function setCurrentFile(fileName : string) {
    currFileNameRef.value = fileName;
    fileContentRef.value = ''; // Fetch file data here ?
  }
  function updateFileContent(payload: string | number) {
    fileContentRef.value = payload.toString();
  };

  async function saveCurrentFile() {
    // Incorporate source in repo data
    const source = "github.com";
    const repoName = currRepoNameRef.value;
    const fileName = currFileNameRef.value;
    const content = fileContentRef.value;
    try {
      const res = await saveFile("github.com", repoName, fileName, content);
    } catch(e) {
      // Display error
      console.log("File not saved");
    }
  }
  async function fetchCurrentFile() {
    const source = "github.com";
    const repoName = currRepoNameRef.value;
    const fileName = currFileNameRef.value;
    const res = await fetchFile(source, repoName, fileName);
    fileContentRef.value = res;
  }
  return { repoName, fileName, fileContent, setCurrentRepo, setCurrentFile, updateFileContent, saveCurrentFile, fetchCurrentFile };
});
