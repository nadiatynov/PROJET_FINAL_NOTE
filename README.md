# PROJET FINAL SOUTIEN #

# PRESENTATION :
Ce projet est une application web conçu pour les fans de pokémon. Nommé MyPokéPack elle permet au uutilisateur de collectionner des cartes pokemon virtuel (via une API) grâce a un système de packs de cartes a ouvrir.

Cette application a été construite a l’aide de pls langages :

- Golang pour le server et le code général 
- HTML & CSS pour les pages web
De plus, une base de données a été utilisée pour stocker de manière durable les informations des utilisateurs et leur information relative aux collections cartes.


# TUTO D’INSTALLATION :
Pour tester ce projet, suivez les étapes suivantes : 
- Ouvrez VSCode et sélectionnez le dossier dans lequel vous souhaitez installer les fichiers
- Ouvrez le terminal intégré et faite la commande "git clone [lien github du projet]"
- Après l'installation des fichiers, faites la commande "go mod tidy" au même endroit, pour vérifier et compléter les dépendances nécessaires
- Puis la commande "go get github.com/mattn/go-sqlite3 " afin d'installer la bibliothèque pour la base de donnée
- Enfin, vous n'avez plus qu'à faire la commande "go run ." et vous rendre sur votre navigateur sur l'adresse "http://localhost:8080"


# POUR UNE MEILLEURE UTILISATION : 
- Ne pas revenir en arrière avec les flèches (cela peut fausser les résultats)
- Si vous souhaiter quitter la page sans vous déconnecter, aller dans inspection, puis supprimez manuellement le user. A la suite de cela, effacer la base de donnée dans VSCode (elle se rechargera automatiquement avec le go run .)
- Si vous avez donner un nom au dossier avant d’avoir fait un git clone, ne pas oublier de faire la commande cd + tab pour se mettre a la racine du projet y compris le repository.


# GUIDE DE L’APPLICATION: 
L’utilisation du site commence par une étape d’authentification. L’utilisateur peut soit s’inscrire en créant un compte avec un nom d’utilisateur, une adresse mail et un mot de passe, soit simplement se connecter s’il possède déjà un compte. Cette étape permet d’associer chaque joueur à un ID unique et de garantir une meilleure gestion de l’attribution des collections de cartes par joueur par exemple. 

Après cette étape, si l’inscription a abouti l’utilisateur est redirigé automatiquement vers la page connexion et seulement a l'issue de celle ci, si elle réussie, que nous somme dirigé automatiquement vers le Dashboard où il pourra observer le nombres total de carte qu’il possède, sa collection de carte et le bouton « ouvrir un pack » qui lui permettra  d’ouvrir un pack de 3 cartes pokémon qui seront automatiquement ajouter à sa collection personnelles de cartes.


# UTILISATION D’IA RESPCTUEUSE CONCERNANT Marjane : 
- L’ia m’a rappeler dans quel situation il est préférable utiliser les pointeurs. Quand j’utilise un Scan par exemple je dois mettre devant les variables un &. Ce pointeurs donne l’adresse mémoire de la variable.(cependant cette notion reste légèrement flou mais la logique global je l’ai comprise) mais grossièrement dans cette exemple par exemple : 
    var c Card
    		rows.Scan(&c.Nom, &c.Type, &c.Image) // IA pointeur
    		cards = append(cards, c)
la variable « c » reçoit les valeurs nom type et image.

- L’ai m’a permi aussi d’apprendre la différence entre « **queryrow** » qui sera utiliser pour une requête d’une seul ligne et « **query** » qui est utiliser pour plusieurs requêtes par exemple ou il y a plusieurs join ou des sous requêtes dans une même requête. Ainsi dans le handler.go j’ai pu expérimenter les 2.

- De plus, elle m’a permis de comprendre la structure (uniquement) d’un handler mêlant librairie random et api en même temps. A la suite de cela, j’ai compris que lorsqu’il faut appeler l’API dans le random et bien il fallait la mettre en Sprintf car sa renvoie une chaine de caractère et il se trouve dans une boucle for car a chaque tour on construit un nouvel URL et on met « %d » pour remplacer ce bout {id_du_pokemon} par un nombre entier. Ainsi quand i est = 2 par exemple sa renvoie le pokémon 2.


- Lors de l’étape de la connexion il fallait vérifier si l’utilisateur avait déjà été connecter ou inscrit (grâce à son ID) pour renvoyer vers la route / si pas de connexion. Ainsi l’ia m’a aiguiller sur le morceau de code ci-dessous qui permet de vérifier si l’utilisateur a un id ou non. J’ai compris avec son aide que cette étape est crucial car par la suite on aura besoin de son ID pour lui afficher sa collection et son nombre de cartes respectifs. Si il y a pas d’id on ne pourra savoir à qui appartient quoi (j’ai eu du mal a mêler le handler et la fonction standard « Verifconnect »). Voici le bout de code :

    userId := Verifconnect(username, mdp)
    	if userId == 0 { //si pas de user ou mauvais identfiant on revoie a P princiapl
    et la fonction standard Verifconnect
  
Le Verifconnect cependant a été coder à l’aide de mes connaissacne et ressources internet. Par ailleurs grâce à un rappel de l’ia sur les pointeurs, j’ai compris automatiquement que dans la fonction Verifconnect l’id devait être en pointeur car on utilise  un Scan et que celui-ci devait lire la valeur venant de la base de donnée et l’écrire dans le handler.go (en gros il devait l’avoir en mémoire).

- Le CSS a été fait avec l’ai 



