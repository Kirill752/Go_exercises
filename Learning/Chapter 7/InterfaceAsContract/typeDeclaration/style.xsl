<?xml version="1.0" encoding="UTF-8"?>
<xsl:stylesheet version="1.0" xmlns:xsl="http://www.w3.org/1999/XSL/Transform">
    <xsl:template match="/">
        <html>
            <head>
                <title>Book Catalog</title>
                <style>
                    body {
                    font-family: Arial, sans-serif;
                    background-color: #f4f4f4;
                    color: #333;
                    }
                    h1 {
                    text-align: center;
                    color: #4CAF50;
                    }
                    .book {
                    background: #fff;
                    margin: 20px;
                    padding: 15px;
                    border-radius: 8px;
                    box-shadow: 0 0 10px rgba(0, 0, 0, 0.1);
                    }
                    .title {
                    font-size: 1.5em;
                    color: #4CAF50;
                    }
                    .author {
                    font-style: italic;
                    color: #555;
                    }
                    .price {
                    font-weight: bold;
                    color: #e91e63;
                    }
                </style>
            </head>
            <body>
                <h1>Book Catalog</h1>
                <xsl:for-each select="catalog/book">
                    <div class="book">
                        <div class="title">
                            <xsl:value-of select="title" />
                        </div>
                        <div class="author">Author: <xsl:value-of select="author" /></div>
                        <div class="genre">Genre: <xsl:value-of select="genre" /></div>
                        <div class="price">Price: $<xsl:value-of select="price" /></div>
                        <div class="description">
                            <xsl:value-of select="description" />
                        </div>
                    </div>
                </xsl:for-each>
            </body>
        </html>
    </xsl:template>
</xsl:stylesheet>