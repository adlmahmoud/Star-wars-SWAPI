document.addEventListener("DOMContentLoaded", () => {
    // 1. Au chargement, on demande au SERVEUR quels sont les favoris
    fetch('/api/favorites')
        .then(response => response.json())
        .then(data => {
            // data est une liste d'objets : [{id: "Luke", type: "people"}, ...]
            data.forEach(fav => {
                // On cherche le bouton qui correspond à ce favori
                // Sélecteur : bouton qui a data-id="Luke" ET data-type="people"
                const btn = document.querySelector(`.fav-btn[data-id="${fav.id}"][data-type="${fav.type}"]`);
                if (btn) {
                    btn.classList.add('active');
                }
            });
        })
        .catch(err => console.error("Erreur chargement favoris:", err));

    // 2. On configure les clics sur les étoiles
    const buttons = document.querySelectorAll('.fav-btn');
    buttons.forEach(btn => {
        btn.addEventListener('click', (e) => {
            e.preventDefault();
            toggleFavorite(btn);
        });
    });
});

function toggleFavorite(btn) {
    const id = btn.dataset.id;
    const type = btn.dataset.type;

    // On envoie la demande au serveur
    fetch('/api/favorites/toggle', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify({ id: id, type: type })
    })
    .then(response => response.json())
    .then(data => {
        // data.added est 'true' si ajouté, 'false' si retiré
        if (data.added) {
            btn.classList.add('active');
            console.log("Ajouté:", id);
            
            // Petite animation
            btn.style.transform = "scale(1.3)";
            setTimeout(() => btn.style.transform = "scale(1)", 200);
        } else {
            btn.classList.remove('active');
            console.log("Retiré:", id);
            
            // Si on est sur la page des favoris, on peut recharger pour voir la carte disparaitre
            if (window.location.pathname === "/favorites") {
                location.reload();
            }
        }
    })
    .catch(err => console.error("Erreur toggle:", err));
}

// Fonction spéciale pour le bouton "Retirer" de la page favoris
function removeFavFromPage(id, type) {
    // On simule un clic sur un bouton
    // Ou on appelle directement l'API
    fetch('/api/favorites/toggle', {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({ id: id, type: type })
    }).then(() => location.reload());
}