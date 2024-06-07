-- Consulta para analizar los patrones de reservas
SELECT
    DAYNAME(FechaReserva) AS DiaDeLaSemana,
    HOUR(HoraReserva) AS HoraDelDia,
    NumeroDeComensales,
    COUNT(*) AS Frecuencia
FROM
    Reservas
GROUP BY
    DiaDeLaSemana,
    HoraDelDia,
    NumeroDeComensales
ORDER BY
    Frecuencia DESC;
