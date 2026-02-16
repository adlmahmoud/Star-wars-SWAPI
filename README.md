# Star Wars Fanpage - Groupie Tracker

Ce projet est une application web développée en **Go** (Golang) dans le cadre du module "Groupie Tracker" à Ynov Campus. L'objectif est de créer un site web dynamique exploitant une API REST pour afficher, rechercher et manipuler des données.

## 🌟 À propos du projet

Cette application est une encyclopédie interactive de l'univers Star Wars. Elle permet aux utilisateurs de naviguer à travers les personnages, films, vaisseaux et planètes de la saga, tout en offrant des fonctionnalités avancées comme la recherche, le filtrage et la gestion de favoris.

L'application respecte l'architecture **MVC** (Model-View-Controller) et n'utilise aucun framework CSS ou JS externe, conformément aux contraintes pédagogiques.

### Thème et Données
* **Thème :** Star Wars (Design sombre et futuriste, accents jaune emblématique).
* **API utilisée :** [SWAPI (The Star Wars API)](https://swapi.dev/)
* **API Complémentaire :** [Star Wars Visual Guide](https://starwars-visualguide.com/) (pour les images).

## 🚀 Fonctionnalités

L'application implémente l'ensemble des fonctionnalités demandées :

* **Exploitation d'API :** Utilisation de plus de 5 endpoints (People, Films, Planets, Species, Starships, Vehicles).
* **FT1 - Recherche :** Système de recherche textuelle permettant de trouver des ressources spécifiques.
* **FT2 - Filtres :** Système de filtres cumulables (ex: filtrer les véhicules par coût, capacité de chargement et longueur).
* **FT3 - Pagination :** Navigation fluide entre les différentes pages de résultats de l'API.
* **FT4 - Favoris Persistants :** Possibilité d'ajouter ou retirer des éléments des favoris. La liste est sauvegardée côté serveur dans un fichier JSON (`favorites.json`), assurant la persistance des données même après redémarrage.

## 🛠️ Stack Technique

* **Langage Backend :** Go (Golang) - Librairie standard (`net/http`, `html/template`, `encoding/json`).
* **Frontend :** HTML5, CSS3 (Grid/Flexbox, Variables CSS, Animations).
* **Architecture :** Arborescence standard Go (cmd, internal, templates, assets).

## 📦 Installation et Lancement

1.  **Cloner le dépôt :**
    ```bash
    git clone [https://github.com/ton-pseudo/swapi-project.git](https://github.com/ton-pseudo/swapi-project.git)
    cd swapi-project
    ```

2.  **Lancer le serveur :**
    Placez-vous dans le dossier contenant le `main.go` (ex: `src/cmd`) :
    ```bash
    go run main.go
    ```

3.  **Accéder au site :**
    Ouvrez votre navigateur et allez sur : `http://localhost:8080` (ou le port configuré).

## 📂 Structure du projet

/src
├── /cmd            # Point d'entrée (main.go)
├── /controllers    # Logique de contrôle (gestion des requêtes HTTP)
├── /models         # Structures de données (mapping JSON)
├── /services       # Logique métier (appels API, filtres)
├── /routers        # Définition des routes
/templates           # Fichiers HTML (Go Templates)
/assets              # CSS, Images, JS
favorites.json       # Fichier de persistance des favoris
## 📝 Gestion de projet (Synthèse)

*La FAQ détaillée concernant l'organisation du travail, la répartition des tâches et la gestion du temps est disponible directement sur le site, dans la page **"À Propos"**, conformément aux consignes.*

---
© 2026 - Projet Étudiant Ynov Campus