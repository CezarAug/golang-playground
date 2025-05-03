CREATE TABLE PERSONS(
    ID SERIAL PRIMARY KEY,
    NAME VARCHAR,
    ABOUT VARCHAR
);

INSERT INTO PERSONS(NAME, ABOUT) VALUES
('Naruto Uzumaki', 'Naruto Uzumaki é um ninja da Vila da Folha que sonha em se tornar Hokage. Apesar de uma infância solitária, ele cresceu determinado e cheio de energia, conquistando respeito e amizade ao longo de sua jornada.'),
('Asuka Langley Soryu', 'Asuka Langley Soryu é uma das pilotos de EVA em Neon Genesis Evangelion. Inteligente e orgulhosa, ela esconde sua vulnerabilidade por trás de uma personalidade forte e determinada. Suas batalhas são tão internas quanto externas.');