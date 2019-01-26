const parser = new DOMParser();

export default function unescapeHTML(encodedString) {
  const DOMNode = parser.parseFromString(encodedString, 'text/html');
  return DOMNode.body.textContent;
}
