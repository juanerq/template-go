# Instrucciones para generar mensajes de commit

Siempre que generes un mensaje de commit para este repositorio, sigue estas
reglas y **responde únicamente con el texto del mensaje**, sin explicaciones.

---

## 1. Formato general (Conventional Commits)

Usa siempre este formato:

<tipo>(<scope>): <resumen breve en español>

- bullet 1
- bullet 2
- bullet 3

### `<tipo>`

Debe ser uno de los siguientes:

- feat → nueva funcionalidad
- fix → corrección de errores
- refactor → mejora interna sin cambio funcional
- docs → documentación
- style → cambios de formato (no altera código de producción)
- test → pruebas
- chore → tareas de soporte (config, build, deps, scripts)

### `<scope>`

- Usa el nombre del módulo, carpeta o microservicio modificado.
- Ejemplos: template-mssql, ms-customer360, api, infra, tests.

### `<resumen breve>`

- En español.
- Tiempo presente / imperativo.
- Máx. ~50 caracteres (el encabezado completo no puede exceder 70 caracteres).
- Sin punto al final.

---

## 2. Cuerpo del commit

Después de la primera línea, agrega una lista de cambios:

- Cada línea debe comenzar con `- `.
- Resume cambios técnicos específicos.
- No repitas literalmente el título.
- Menciona si hubo:
  - actualización de DTOs,
  - ajustes en configuraciones (env, Docker, Husky),
  - cambios en integración o estructura del proyecto,
  - modificaciones en pruebas.

Ejemplo de bullets:

- actualiza validación de campos en DTO AppRequest
- refactoriza OracleService para mejor manejo de errores
- ajusta configuración de Husky para lint-staged

---

## 3. Estilo

- No usar emojis.
- Español neutro.
- No usar frases genéricas (“varios cambios”, “arreglos”).
- Cada commit debe representar un cambio coherente.

---

## 4. Ejemplos

### Ejemplo 1

feat(api): agrega validación de parámetros

- valida que los placeholders reciban valores requeridos
- agrega manejo de error cuando falta un parámetro
- actualiza pruebas unitarias de OracleService

### Ejemplo 2

fix(infra): corrige autenticación Salesforce

- captura excepciones de autenticación y mapea códigos 401/403
- agrega logs estructurados con correlationId
- actualiza documentación de guía de integración
