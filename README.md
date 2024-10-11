# Crucigrama en Golang
TUI de crucigrama en Golang consumiendo la API de Wikipedia para el curso de Programación Para Competición Avanzado.
Incluye el efecto matrix para el crucigrama.

## Capturas de pantalla
### Proceso (GIF)
![crucigrama](https://github.com/user-attachments/assets/fc427122-4204-4421-8223-3733d7f0bb9b)

### Llamada a la app
![image](https://github.com/user-attachments/assets/880ba651-66f3-49ad-a645-5292d7f2c850)

### Entrar un tema
![image](https://github.com/user-attachments/assets/8575d828-a427-4966-a91b-68c5da8bd5fa)

### Ver artículos relacionados
![image](https://github.com/user-attachments/assets/20506846-b870-4534-b28a-b51744b1a29c)

### Filtrar artículo de interés
![image](https://github.com/user-attachments/assets/b61a51d0-8129-4ae7-87f9-3869bf3853b2)

### Ingresar tamaño y número de palabras
![image](https://github.com/user-attachments/assets/be282276-71ee-4e2b-ab48-5175049ffb47)

### Crucigrama generado
![image](https://github.com/user-attachments/assets/73b449f2-1e69-40e0-bb68-adfd97278bd7)

### Efecto matrix generado
![image](https://github.com/user-attachments/assets/69de990d-4370-47c3-b46b-3cc20c458085)

## Modo de Uso
- Ingresa cualquier tema
- El programa buscará articulos relacionados en Wikipedia
- Filtra y escoge un artículo
- Ingresa el tamaño del crucigrama y la cantidad de palabras que se usarán
- El sistema generará un crucigrama con ese artículo y lo mostrará
- Presiona `m` para ver un efecto matrix con las letras del crucigrama

## Módulos
- `core`: Los algoritmos para generar el crucigrama y la matriz
- `tui`: Interfaz de la aplicación
- `stopwords`: Lista de stopwords a quitar de los artículos
- `wikipedia`: Funciones para hacer llamadas a la API de wikipedia
