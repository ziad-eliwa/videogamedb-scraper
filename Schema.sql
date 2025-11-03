CREATE DATABASE VideoGames;
USE VideoGames;

CREATE TABLE IF NOT EXISTS Game (
    Name VARCHAR(255) PRIMARY KEY,
	CoverURL TEXT,
    Description TEXT
);

CREATE TABLE IF NOT EXISTS People (
	MobyID INTEGER PRIMARY KEY,
    Name TEXT
);

CREATE TABLE IF NOT EXISTS GameDirector (
	GameName VARCHAR(255),
	DirectorID INTEGER,
    DirectorRole VARCHAR(255),
    PRIMARY KEY(GameName,DirectorID,DirectorRole),
    FOREIGN KEY (GameName) REFERENCES Game(Name)
    ON UPDATE CASCADE,
    FOREIGN KEY (DirectorID) REFERENCES People(MobyID)
    ON UPDATE CASCADE
);


CREATE TABLE IF NOT EXISTS Category ( -- done
    CategoryName VARCHAR(100),
    CategoryType VARCHAR(100),
    Description TEXT,
    PRIMARY KEY (CategoryName, CategoryType)
);

CREATE TABLE IF NOT EXISTS GameCategory ( 
    GameName VARCHAR(255),
    CategoryName VARCHAR(100),
    CategoryType VARCHAR(100),
    PRIMARY KEY (GameName, CategoryName, CategoryType),
    FOREIGN KEY (GameName) REFERENCES Game(Name) ON UPDATE CASCADE,
    FOREIGN KEY (CategoryName, CategoryType) REFERENCES Category(CategoryName, CategoryType) 
    ON UPDATE CASCADE
);

CREATE TABLE IF NOT EXISTS User (-- faker
    Username VARCHAR(50) PRIMARY KEY,
    Birthdate DATE,
    Gender VARCHAR(20),
    Country VARCHAR(100)
);

CREATE TABLE IF NOT EXISTS UserEmails ( -- faker
    Username VARCHAR(50),
    EmailAddress VARCHAR(255),
    PRIMARY KEY (Username, EmailAddress),
    FOREIGN KEY (Username) REFERENCES User(Username) ON UPDATE CASCADE
);

CREATE TABLE IF NOT EXISTS Collection ( -- faker
    Username VARCHAR(50) NOT NULL,
    CollectionName VARCHAR(255) PRIMARY KEY,
    Description TEXT,
    Privacy VARCHAR(7) CHECK (Privacy IN ('public', 'private')),
    FOREIGN KEY (Username) REFERENCES User(Username) ON UPDATE CASCADE
);

CREATE TABLE IF NOT EXISTS CollectionGames ( -- faker
    CollectionName VARCHAR(255),
    GameName VARCHAR(255),
    PRIMARY KEY (CollectionName, GameName),
    FOREIGN KEY (CollectionName) REFERENCES Collection(CollectionName) ON UPDATE CASCADE,
    FOREIGN KEY (GameName) REFERENCES Game(Name) ON UPDATE CASCADE
);

CREATE TABLE IF NOT EXISTS Platforms ( -- done
    Name VARCHAR(100) PRIMARY KEY,
    Description TEXT,
    Years_Active TEXT
);

CREATE TABLE IF NOT EXISTS GameRelease (
	GameName VARCHAR(255),
    PlatformName VARCHAR(100),
    ReleaseDate DATE,
    ReleaseComment TEXT,    
    RetailPrice DECIMAL,
    PRIMARY KEY (GameName,PlatformName, ReleaseDate),
    FOREIGN KEY (GameName) REFERENCES Game(Name) ON UPDATE CASCADE,
    FOREIGN KEY (PlatformName) REFERENCES Platforms(Name) ON UPDATE CASCADE
);

CREATE TABLE IF NOT EXISTS TechSpecs (
    SpecName VARCHAR(100),
    SpecAttribute VARCHAR(100),
    PRIMARY KEY (SpecName, SpecAttribute)
);

CREATE TABLE IF NOT EXISTS TechSpecRelease (
    GameName VARCHAR(255),
    PlatformName VARCHAR(100),
    SpecName VARCHAR(100),
    SpecAttribute VARCHAR(100),
    PRIMARY KEY (GameName,PlatformName, SpecName, SpecAttribute),
    FOREIGN KEY (GameName,PlatformName) REFERENCES GameRelease(GameName,PlatformName) 
    ON UPDATE CASCADE,
    FOREIGN KEY (SpecName, SpecAttribute) REFERENCES TechSpecs(SpecName, SpecAttribute) 
    ON UPDATE CASCADE
);

CREATE TABLE IF NOT EXISTS MaturityRating (
    MaturityRatingType VARCHAR(50),
    MaturityRatingAttribute VARCHAR(100),
    PRIMARY KEY (MaturityRatingType, MaturityRatingAttribute)
);

CREATE TABLE IF NOT EXISTS MaturityRatingRelease (
    GameName VARCHAR(255),
    PlatformName VARCHAR(100),
    MaturityRatingType VARCHAR(50),
    MaturityRatingAttribute VARCHAR(100),
    PRIMARY KEY (GameName,PlatformName, MaturityRatingType, MaturityRatingAttribute),
    FOREIGN KEY (GameName,PlatformName) REFERENCES GameRelease(GameName,PlatformName) 
    ON UPDATE CASCADE,
    FOREIGN KEY (MaturityRatingType, MaturityRatingAttribute) REFERENCES
    MaturityRating(MaturityRatingType, MaturityRatingAttribute) ON UPDATE CASCADE
);

CREATE TABLE IF NOT EXISTS UserPlatform (
    Username VARCHAR(50),
    PlatformName VARCHAR(100),
    PRIMARY KEY (Username, PlatformName),
    FOREIGN KEY (Username) REFERENCES User(Username) ON UPDATE CASCADE,
    FOREIGN KEY (PlatformName) REFERENCES Platforms(Name) ON UPDATE CASCADE
);

CREATE TABLE IF NOT EXISTS Company (
    Name VARCHAR(255) PRIMARY KEY,
    Overview TEXT,
    LogoURL VARCHAR(255),
    YearStarted YEAR,
	YearEnded YEAR,
    Country VARCHAR(100)
);

CREATE TABLE IF NOT EXISTS Developer (
    DeveloperName VARCHAR(255),
    GameName VARCHAR(255),
    PRIMARY KEY (DeveloperName, GameName),
    FOREIGN KEY (DeveloperName) REFERENCES Company(Name) ON UPDATE CASCADE,
    FOREIGN KEY (GameName) REFERENCES Game(Name) ON UPDATE CASCADE
);

CREATE TABLE IF NOT EXISTS Publisher (
    PublisherName VARCHAR(255),
    GameName VARCHAR(255),
    PRIMARY KEY (PublisherName, GameName),
    FOREIGN KEY (PublisherName) REFERENCES Company(Name) ON UPDATE CASCADE,
    FOREIGN KEY (GameName) REFERENCES Game(Name) ON UPDATE CASCADE
);


CREATE TABLE IF NOT EXISTS UserReview (
    Rating DECIMAL(2,1) NOT NULL,
    ReviewComment TEXT,
    Username VARCHAR(50) NOT NULL,
    GameName VARCHAR(255) NOT NULL,
    PlatformName VARCHAR(255) NOT NULL,
    ReviewDate DATE,
    PRIMARY KEY(GameName,Username),
    FOREIGN KEY (Username) REFERENCES User(Username) ON UPDATE CASCADE,
    FOREIGN KEY (GameName,PlatformName) REFERENCES GameRelease(GameName,PlatformName) ON UPDATE CASCADE
);

CREATE TABLE IF NOT EXISTS MobyGameReview (
	GameName VARCHAR(255),
    PlatformName VARCHAR(255),
	MobyScorePlayers DECIMAL(2,1),
    MobyScoreCritic DECIMAL(2,1),
    NumberOfPlayers INTEGER,
	NumberOfCritics INTEGER,
    MobyRating DECIMAL(2,1),
    PRIMARY KEY(GameName,PlatformName),
	FOREIGN KEY (GameName) REFERENCES Game(Name) ON UPDATE CASCADE,
    FOREIGN KEY (PlatformName) REFERENCES Platforms(Name) ON UPDATE CASCADE
);