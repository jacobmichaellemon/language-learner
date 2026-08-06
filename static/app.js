// JS: loop through "a few elements"
const nodes = document.querySelectorAll(".js-item");

nodes.forEach((node, i) => {
if (i >= nodes.length) return;
    node.addEventListener("click", () => {
        const value = node.dataset.value;
        // Copy the text inside the text field
        navigator.clipboard.writeText(value)
        showToast("Copied to clipboard")
    });
});

function showToast(msg) {
    const t = document.createElement('div');
    t.textContent = msg;

    t.style.position = 'fixed';
    t.style.left = '50%';
    t.style.bottom = '24px';
    t.style.transform = 'translateX(-50%)';
    t.style.padding = '8px 12px';
    t.style.background = 'rgba(0,0,0,0.75)';
    t.style.color = '#fff';
    t.style.fontSize = '12px';
    t.style.borderRadius = '999px';
    t.style.zIndex = 9999;
    t.style.opacity = '0';
    t.style.transition = 'opacity 160ms ease, transform 160ms ease';

    document.body.appendChild(t);

    requestAnimationFrame(() => {
      t.style.opacity = '1';
      t.style.transform = 'translateX(-50%) translateY(0)';
    });

    setTimeout(() => {
      t.style.opacity = '0';
      t.remove();
    }, 900);
  }