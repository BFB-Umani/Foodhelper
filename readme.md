# Foodhelper

Foodhelper är ett CLI program skrivet i Go för att hämta recept som passar din kyl.
Tanken är att du med hjälp av ett gäng olika kategorier kan få ut recept som passar ditt sug eller tillochmed vad du har hemma i kylen!

Programmet använder sig av [Huh?](https://github.com/charmbracelet/huh) från [charm bracelet](charm.sh) för formulärhantering i command line där du lätt kan fylla i vad du är sugen på idag.

För att installera programmet till din CLI behöver du enbart klona ner projektet och sen köra "go install". Om du istället vill skapa en körbar fil så använd istället kommandot "go build" vilket kommer skapa en executeable i roten av projektet, kör sen denna fil i command line.

Projektet kommer laddat med en massa recept som täcker de flesta kategorier, men om du vill skapa recept så går det fint att göra i programmet också vilket kommer skapa 2 nya mappar antingen i samma mapp som du har din körbara fil **eller** i mappen du står i när du kör programmet i det fall att du användt "go install" för att installera programmet.

För att sen kunna få fram dina egna recept så måste du köra programmet i samma mapp som du skapat dina recept!

Et lättare program för att komma igång med att koda i Go samt för att lära mig skriva CLI program med [Charm](charm.sh).