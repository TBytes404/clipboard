import "htmx.org";

async function paste(ev: MouseEvent) {
  (ev.currentTarget as HTMLButtonElement).value =
    await navigator.clipboard.readText();
}

async function copy(text: string) {
  await navigator.clipboard.writeText(text);
}

document.getElementById("adder")!.onclick = paste;

const toast = document.getElementById("toast") as HTMLElement;

Array.from(document.getElementsByClassName("blob")).forEach((el) => {
  (el as HTMLDivElement).onclick = () => {
    copy(
      (el.getElementsByTagName("p")[0] as HTMLParagraphElement).textContent!,
    );

    toast.classList.replace("hidden", "flex");
    setTimeout(() => {
      toast.classList.replace("flex", "hidden");
    }, 3000);
  };
});
