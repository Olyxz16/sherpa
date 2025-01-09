type FetchFileResponse = {
    "content": string
}

export async function saveFile(source: string, repoName: string, fileName: string, content: string): Promise<boolean> {
  if(repoName === "") {
    throw "Repository name cannot be empty"
  }
  if(fileName === "") {
    throw "File name cannot be empty"
  }
  const url = "/file";
  const bodyObj = {
    source: source,
    repoName: repoName,
    fileName: fileName,
    content: content
  };
  const headers = new Headers({'content-type': 'application/json'});
  try {
    const resp = await fetch(url, {
      method: "POST",
      headers: headers,
      body: JSON.stringify(bodyObj)
    });
    if(!resp.ok) {
      return false
    }
  } catch(e) {
    throw e
  }
  return true;
}

export async function fetchFile(source: string, repoName: string, fileName: string): Promise<string> {
  const url = "/file?";
  const param = new URLSearchParams({
    source: source,
    repo: repoName,
    file: fileName
  });
  const query = url + param.toString();
  let content = "";
  try {
    const resp = await fetch(query);
    if(!resp.ok) {
      return ""
    }
    const jsonStr = await resp.json() as string
    const json = JSON.parse(jsonStr) as FetchFileResponse
    content = json.content;
  } catch(e) {
    return ""
  }
  return content;
}
